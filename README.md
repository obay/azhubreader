# Azure Event Hub Reader

This Go application reads events from an Azure Event Hub and writes specific events to JSON files. It filters events based on predefined categories and operations, making it useful for monitoring and auditing purposes.

This command line was created specificly to monitor for users and devices creation/deletion in Entra ID. You will find a PowerShell script that also can be used to trigger any other actions based when a new JSON file (new event) is created.

## Installation

You can install azhubreader on any MacOS or Linux machine with Homebrew installed.

You can install Homebrew using the following command:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Once Homebrew is installed, run the following commands to install azhubreader:

```bash
brew tap obay/tap && brew install azhubreader
```

Confirm the tool is installed by running:

```bash
azhubreader version
```

## Usage

Replace the placeholders with your own values and run the following command:

```bash
eventhub_name="EntraID-eh"
consumer_group="EntraID-cg"
eventhub_namespace="EntraID-ns"
EventHubOutput="EventHubOutput"
key_name="RootManageSharedAccessKey"
key="EBzdTLpH8l+4H7xxxIq5rbFpbOVQNuO/f+AExPxeKGo=" # This is a fake key. Replace with your own key.

azhubreader -hub=$eventhub_name -group=$consumer_group -output=$EventHubOutput -namespace=$eventhub_namespace -keyname=$key_name -key=$key
```

### Parameters

- `-hub`: Event Hub name
- `-group`: Consumer group name
- `-output`: Output directory for JSON files
- `-namespace`: Event Hub namespace
- `-keyname`: Shared Access Key name
- `-key`: Shared Access Key

All parameters are required.

## Configuration

The application is configured to monitor the following:

- Categories: "AuditLogs"
- Operations: "Add user", "Delete user"

To modify these, update the `monitoredCategories` and `monitoredOperations` slices in the `ReadEventHub` function.

## Output

The application will create JSON files in the specified output directory. Each file will be named using the format:
YYYY-MM-DD_HH-MM-SS-Operation_Name.json

## Notes

- The application will continue running until manually stopped.
- Ensure you have the necessary permissions to read from the Event Hub and write to the output directory.
