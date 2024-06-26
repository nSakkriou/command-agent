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
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
            background-color: #f4f4f4;
        }

        h1, h2, h3 {
            margin: 0;
            padding: 10px 0;
        }

        h1 {
            border-bottom: 2px solid #333;
        }

        .blue-bg {
            background-color: #bde0fe;
            padding: 10px;
            border-radius: 5px;
        }

        .p-2 {
            padding: 10px;
        }

        .mb-2 {
            margin-bottom: 20px;
        }

        .command {
            margin-bottom: 20px;
            padding: 10px;
            border: 1px solid #ddd;
            border-radius: 5px;
            background-color: #fff;
        }

        .command h2 {
            margin-bottom: 5px;
        }

        .command p {
            margin: 5px 0;
        }

        .command button {
            margin: 5px 0;
            padding: 10px 15px;
            background-color: #007bff;
            color: white;
            border: none;
            border-radius: 5px;
            cursor: pointer;
        }

        .command button:hover {
            background-color: #0056b3;
        }

        .command hr {
            border: 0;
            border-top: 1px solid #eee;
            margin: 10px 0;
        }

        .log-container {
            background-color: #fff;
            border: 1px solid #ddd;
            border-radius: 5px;
            padding: 10px;
        }

        pre {
            background-color: #bde0fe;
            padding: 10px;
            border-radius: 5px;
            overflow-x: auto;
        }

        .row {
            display: flex;
            align-items: center;
            margin-bottom: 10px;
        }

        .row h2 {
            flex: 1;
        }

        .row button {
            margin-left: 10px;
        }

        .success {
            color: white;
            background-color: #a7c957;
            padding: 5px;
            border-radius: 5px;
        }

        .error {
            color: white;
            background-color: #ef233c;
            padding: 5px;
            border-radius: 5px;
        }

        .pointer {
            cursor: pointer;
        }

        .mr-2 {
            margin-right : 10px;
        }

        .mt-0 {
            margin-top : 0;
        }

        .col {
            display: flex;
            flex-direction: column;
        }

        #loadingScreen {
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background-color: rgba(255, 255, 255, 0.4);
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            z-index: 1000;
        }

        .spinner {
            border: 4px solid #f3f3f3;
            border-radius: 50%;
            border-top: 4px solid #3498db;
            width: 40px;
            height: 40px;
            animation: spin 1s linear infinite;
            opacity: 1;
        }

        @keyframes spin {
            0% { transform: rotate(0deg); }
            100% { transform: rotate(360deg); }
        }
    </style>
</head>

<body>
    <div id="loadingScreen" style="display: none">
        <div class="spinner"></div>
        <p>Loading...</p>
    </div>

    <div id="log-wrapper mb-2">
        <div class="row">
            <h2 class="p-0 mr-2">Logs</h2>
            <button class="blue-bg p-2 pointer" style="border: 0.5px solid black;" id="clear-log">Clear</button>
        </div>

        <div class="row" style="margin-top: 5px;">
            <h3>Response <p id="status-code" style="margin-right: 20px; margin-top: 0px;margin : 0;"></p></h3>
        </div>

        <div class="row" id="success-row">

        </div>

        <pre id="log" class="blue-bg" style="width: 100%; min-height: 200px; overflow-x: auto;margin-right: 20px;"></pre>
    </div>

    <h1 class="mb-2">Endpoint available</h1>
    
	{{ range .EndCommands }}
	<div class="command">
        <div style="display : flex; justify-content: space-between; align-items: center;">
            <div class="col">
                <h2>{{.EndpointName}}</h2>
                <p>{{.Method}}</p>
            </div>

            <div class="col">
                <button class="show-content" data-endpoint="{{.EndpointName}}/content" class="blue-bg p-2 pointer" style="border: 0.5px solid black;">
                    Show script content
                </button>

                <button class="exec-btn" data-method="{{.Method}}" data-endpoint="{{.EndpointName}}" class="blue-bg p-2 pointer" style="border: 0.5px solid black;">
                    Execute
                </button>
            </div>
        </div>
    </div>
	{{ end }}

    <script>
        const log = document.getElementById("log")
        const status = document.getElementById("status-code")
        const successRow = document.getElementById("success-row")
        const loader = document.getElementById("loadingScreen")

        const setSucessRow = (outputs) => {
            cleanContainer("#success-row")

            outputs.forEach(e => {
                console.log(e)
                let div = document.createElement("p")
                div.classList.add("mr-2")
                div.classList.add("mt-0")
                setSuccess(e, div)
                successRow.appendChild(div)
            })
        }

        const setSuccess = (output, elem) => {
            if(output.Success) {
                elem.classList.add("success")
                elem.classList.remove("error")

                elem.textContent = output.Filename + " : Success !"
            }
            else {
                elem.classList.add("error")
                elem.classList.remove("success")

                elem.textContent = output.Filename + " : Failed !"
            }
        }

        document.getElementById("clear-log").addEventListener("click", e => {
            clean()
        })

        document.querySelectorAll(".show-content").forEach(e => {
            const endpoint = e.dataset.endpoint

            e.addEventListener("click", event => {
                loader.style.display = "flex"
                clean()

                fetch(endpoint)
                .then(response => {
                    return response.text()
                })
                .then(response => {
                    log.textContent = response
                    loader.style.display = "none"
                })
                .catch(error => log.textContent = "Failed to fetch data : " + error)  
        })})

        document.querySelectorAll(".exec-btn").forEach(e => {
            const endpoint = e.dataset.endpoint
            const method = e.dataset.method

            e.addEventListener("click", event => {
                loader.style.display = "flex"
                clean()
                
                fetch(endpoint, {
                    method : method
                })
                .then(response => {
                    status.textContent = "Status code : " + response.status
                    return response.json()
                })
                .then(res => {
                    setSucessRow(res.Outputs)
                    log.textContent = JSON.stringify(res, null, 2).replace(/\\n/g, '\n\t\t');
                    loader.style.display = "none"
                    
                })
                .catch(error => log.textContent = "Failed to fetch data : " + error)  
        })})

        const cleanContainer = (tag) => {
            var div = document.querySelector(tag);
            while(div.firstChild){
                div.removeChild(div.firstChild);
            }
        }

        const clean = () => {
            status.textContent = ""
            log.textContent = ""
            cleanContainer("#success-row")
        }
    </script>
</body>

</html>
	`
}
