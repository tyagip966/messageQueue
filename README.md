# Message Queuing Service

### Having Below Components

1. Queuing Service (manage connections of consumer and producers, manage data queue)
2. Producer (Produce data which will store inside Queuing Service)
3. Consumer (Listen/Consume data from Queuing Service)

### TODO
1. clone the repository
```
  git clone https://github.com/tyagip966/messageQueue.git
```

### Queuing Service

Below are the instructions to run this

```
1. cd queuingService/
2. go run server.go
```

### Consumer
Open another terminal for Consumer (you'll run multiple consumers in multiple terminal windows)
```
1. cd consumer
2. go run consumer.go
```
Now you'll get your output in the terminal screen whenever producer will produce some data

### Producer
Open another terminal For Producer (you'll run multiple producers in multiple terminal windows)
```
1. cd producer
2. go run main.go
```
Now you'll put your input json inputs in the terminal screen and on consumers screen you'll get the expected output.

#### We have four windows in below snap shot

1. Server (queuingService) (upper left)
2. Producer  (upper right)
3. Consumer 1 (bottom left)
4. Consumer 2 (bottom right)

### Populated three json docs and in the ouput you'll get timestamp as an extra field of json doc.

![image](https://user-images.githubusercontent.com/29175591/143047250-87451293-84e7-4b64-b5f9-fa519193a563.png)
