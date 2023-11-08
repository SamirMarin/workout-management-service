# workout-management-service
Service that manages the creation and querying of workouts

## Running the service locally
1. Start local running dynamodb
```bash
# must be running docker to run this command
docker-compose up
```
2. Build the service
```bash
go build -o workout-management-service
```

3. Run the service
```bash
./workout-management-service
```

## Making requests to the service locally
1. Create a workout
```bash
curl -X POST http://localhost:1323/create \
-H "Content-Type: application/json" \
-d '{
  "owner": "Samir",
  "name": "Run The Interval",
  "category": "running",
  "equipment": {
    "name": "Running Equipment",
    "description": "If indoors we need a threadmill, if outdoors a good place to run fast for 3 min without interruption"
  },
  "exercises": [
    {
      "name": "Warmup",
      "description": "20min jog"
    },
    {
      "name": "3 min by 5 interval",
      "description": "3min x5 at 5k pace with 1min jog"
    },
    {
      "name": "Cooldown",
      "description": "20min cool down"
    }
  ]
}'

```
2. Get a workout
```bash
curl -X POST http://localhost:1323/get \
-H "Content-Type: application/json" \
-d '{
  "owner": "Samir",
  "name": "Run The Interval"
}'
```