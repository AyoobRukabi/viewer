## Cars Viewer

## Folder structure
```text
Viewer/
├── cmd/
│   ├── server/         # This will hold our main.go file. It will be the entry point wher the application starts
│   
├── internal/           
│   ├── models/         # This will store the Structs. E.g. what a Car looks like
│   └── service/        # This will handle the Logic. Specifically the code that fetches data from the external "CarsAPI"
|   └── handlers/       # This hanldes HTTP Requests
├── web/                
|   ├── templates/      # This is for HTML files
│   └── service/        # This is for CSS and JavaScript files

```
