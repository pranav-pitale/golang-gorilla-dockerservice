# RESTful API Example with golang gorilla with docker
This is restful api is build using only with **gorilla/mux**.
Data from the file is loaded as the server starts.
Data is stored in In-memory data store.

## Install and Run
```shell
$ install docker
$ docker build ./
$ docker images
$ docker run -i -t -p 3000:3000 [image id]
```

## API Endpoint
- http://localhost:3000/api/gettotalvelociraptor/{timestamp}
   - `GET`: get total number of velociraptor in +/- 3years of timestamp,where timestamp is ISO datetime viz. 2014-10-08T19:02:17-08:00
- http://localhost:3000/api/updatevelociraptor
    - `POST`: Updates total number of velociraptor with timestamp or nearest timestamp

## Request Payload for Post updatevelociraptor
```json
{
    "TimeStamp":"2014-11-01T19:02:17-08:00"
    ,"TotalVelociraptor": 40
}
```

## Code struture and logic
```shell
- `main.go` entry point of the web api and also has route congfiguration
- `utils`
    - `dateutil.go`: It has the util functions in order to generate date ranges and check certain time stamp 
                    lies between start and end date
- `hander`
    - `handler.go` : It serves request and returns back the result.
- `models`
    - `velociraptor.go` : It has the model for data store. Year store is a map which holds pointer to month store by year                         key.Month store is a map which holds pointer to list of velociraptor by month key of each time                         stamp of month for year in file. Velociraptor has attribute time and total.
- `datastore`
    - `load.go` : It reads the file from data folder and creates data store which is then accessed by api
-`data` : It has csv file of time line

```
