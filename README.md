# Command Agent

Command agent it's go tool who start web api, and link endpoint with shell script.

To start agent you have to create a AgentFile.

It's a JSON file with some mandatory parameter :

```json
{
  "port" : 7890,
  "scripts_folder_path" : "/scripts/",
  "endcommands" : [
    {
      "endpoint_name" : "start",
      "scripts_files_names" : ["start.sh"],
      "method" : "POST"
    }
  ]
}
```

Minimal AgentFile


$ go run main.go
Start process