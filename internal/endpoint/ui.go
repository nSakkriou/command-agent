package endpoint

import (
	"html/template"
	"net/http"

	"github.com/nSakkriou/utils/pkg/agent"
)

func GetUIEndpoint(config *agent.AgentFile) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		t, err := template.New("webpage").Parse(getUIHTMLTemplate())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")

		err = t.Execute(w, config)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

// Mandatory template
func getUIHTMLTemplate() string {
	return `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Command Agent</title>

    <style>
        h1,h2,h3 {
            margin: 0;
        }

        .blue-bg {
            background-color: #bde0fe;
        }

        .p-0 {
            padding: 0;
        }

        .p-2 {
            padding: 10px;
        }

        .m-0 {
            margin: 0;
        }

        .command {
            display: flex;
            justify-content: space-between;
            align-items: center;
        }

        .col {
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
        }

        .row {
            display: flex;
            flex-direction: row;
        }

        .pointer {
            cursor: pointer;
        }

        .mr-2 {
            margin-right: 10px;
        }

        .mb-2 {
            margin-bottom: 10px;
        }
    </style>
</head>

<body>
    <div id="log-wrapper">
        <div class="row">
            <h2 class="p-0 mr-2">Logs</h2>
            <button class="blue-bg p-2 pointer" style="border: 0.5px solid black;" id="clear-log">Clear</button>
        </div>
        <pre id="log" class="blue-bg" style="width: 100vw; min-height: 200px; overflow-x: auto;"></pre>
    </div>

    <h1 class="mb-2">Endpoint available</h1>
    
	{{ range .EndCommands }}
	<div class="command">
        <div class="col">
            <h2>{{.EndpointName}}</h2>
			<p>{{.Method}}</p>
        </div>

        <button class="exec-btn" data-method="{{.Method}}" data-endpoint="{{.EndpointName}}" class="blue-bg p-2 pointer" style="border: 0.5px solid black;">
            Execute
        </button>
    </div>
	{{ end }}

    <script>
        const log = document.getElementById("log")

        document.getElementById("clear-log").addEventListener("click", e => {
            log.innerHTML = ""
        })

        document.querySelectorAll(".exec-btn").forEach(e => {
            const endpoint = e.dataset.endpoint
            const method = e.dataset.method

            e.addEventListener("click", event => {
                fetch(endpoint, {
                    method : method
                })
                .then(response => response.json())
                .then(res => {
                    console.log(res)
                    log.innerHTML = JSON.stringify(res, undefined, 2);
                })
            })
        })
    </script>
</body>

</html>
	`
}
