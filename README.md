# Avregex
Handling SITA and IATA messages using regex in Go. Just for fun 😊

Inspired by [eigthyknots](https://gist.github.com/eightyknots/4372d1166a192d5e9754)

Thinking about this, the task is much larger than i thought. Lets start with the variances of movement messages. For most, they differenciate with the first two letters of line three. But also sometimes some letters 
### MVT Messages
- Departure
- Arrival
- Delays
  - With new time, ED
  - With infinite delay, NI
- Delayed Takeoff
  - Without ETA
  - With ETA
- Return to Ramp
- Return from airborne
- Revised ETA
- Arrival Taxi Time Variance
