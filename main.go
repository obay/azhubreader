package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	eventhub "github.com/Azure/azure-event-hubs-go"
)

func ReadEventHub(eventHubName, consumerGroup, outputDirectory, eventHubNamespace, sharedAccessKeyName, sharedAccessKey string) {
	eventHubConnectionString := fmt.Sprintf("Endpoint=sb://%s.servicebus.windows.net/;SharedAccessKeyName=%s;SharedAccessKey=%s;EntityPath=%s",
		eventHubNamespace, sharedAccessKeyName, sharedAccessKey, eventHubName)

	fmt.Println("Reading events from Event Hub and writing to JSON files...")
	os.MkdirAll(outputDirectory, os.ModePerm)

	hub, err := eventhub.NewHubFromConnectionString(eventHubConnectionString)
	if err != nil {
		fmt.Printf("Error creating Event Hub client: %v\n", err)
		return
	}
	defer hub.Close(context.Background())

	monitoredCategories := []string{"AuditLogs"}
	monitoredOperations := []string{"Add user", "Delete user"}

	ctx := context.Background()
	handler := func(ctx context.Context, event *eventhub.Event) error {
		var jsonNode map[string]interface{}
		if err := json.Unmarshal(event.Data, &jsonNode); err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return nil
		}

		if records, ok := jsonNode["records"].([]interface{}); ok {
			for _, record := range records {
				if r, ok := record.(map[string]interface{}); ok {
					category, _ := r["category"].(string)
					operationName, _ := r["operationName"].(string)
					date, _ := r["time"].(string)

					if category != "" && operationName != "" && date != "" &&
						contains(monitoredCategories, category) &&
						contains(monitoredOperations, operationName) {

						formattedDate := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(date, ":", "-"), "T", "_"), "Z", "")
						fileName := fmt.Sprintf("%s-%s.json", formattedDate, strings.ReplaceAll(operationName, " ", "_"))
						filePath := filepath.Join(outputDirectory, fileName)

						jsonString, err := json.MarshalIndent(r, "", "  ")
						if err != nil {
							fmt.Printf("Error marshaling JSON: %v\n", err)
							continue
						}

						if err := os.WriteFile(filePath, jsonString, 0644); err != nil {
							fmt.Printf("Error writing file: %v\n", err)
							continue
						}

						fmt.Printf("Written event to file: %s\n", fileName)
					}
				}
			}
		}

		return nil
	}

	runtimeInfo, err := hub.GetRuntimeInformation(ctx)
	if err != nil {
		fmt.Printf("Error getting runtime information: %v\n", err)
		return
	}

	for _, partitionID := range runtimeInfo.PartitionIDs {
		_, err := hub.Receive(ctx, partitionID, handler, eventhub.ReceiveWithConsumerGroup(consumerGroup))
		if err != nil {
			fmt.Printf("Error setting up receiver for partition %s: %v\n", partitionID, err)
		}
	}

	// Keep the program running
	<-make(chan struct{})
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	eventHubName := flag.String("hub", "", "Event Hub name")
	consumerGroup := flag.String("group", "", "Consumer group name")
	outputDirectory := flag.String("output", "", "Output directory for JSON files")
	eventHubNamespace := flag.String("namespace", "", "Event Hub namespace")
	sharedAccessKeyName := flag.String("keyname", "", "Shared Access Key name")
	sharedAccessKey := flag.String("key", "", "Shared Access Key")

	flag.Parse()

	// Check if all required parameters are provided
	if *eventHubName == "" || *consumerGroup == "" || *outputDirectory == "" ||
		*eventHubNamespace == "" || *sharedAccessKeyName == "" || *sharedAccessKey == "" {
		fmt.Println("All parameters are required.")
		flag.Usage()
		os.Exit(1)
	}

	ReadEventHub(*eventHubName, *consumerGroup, *outputDirectory, *eventHubNamespace, *sharedAccessKeyName, *sharedAccessKey)
}
