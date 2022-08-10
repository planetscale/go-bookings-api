# Go Booking API

## DB Model

```
hotels
    id
    name
    address
    stars
rooms
    id
    type
    price
    hotel_id
bookings
    id
    arrival
    departure
    guest_id
    room_id
guests
    id
    name
    phone
```

## Used by

The following content uses this demo application:

- [Building a multi-stage pipeline with PlanetScale and AWS](https://planetscale.com/blog/building-a-multi-stage-pipeline-with-planetscale-and-aws)
