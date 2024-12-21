## MVT Message Types

### Departure
```
MVT
SD200/21.PMDFG.CDG
AD1100/1115 EA1500 FRA
DL72/0015
PX112
SI DEICING
```
All messages can also include the date in the departure time fields. ie.
```
MVT
SD200/21.PMDFG.CDG
AD211100/211115 EA1500 FRA
DL72/0015
PX112
SI DEICING
```

### Arrival
```
MVT
SD200/21.PMDFG.CDG
AA1515/1520
```
### Delay
#### With ETD
```
MVT
SD200/21.PMDFG.CDG
ED211125
DL72/0025
```
#### With indefinite delay
Time is set to when new information is expected.
```
MVT
SD200/21.PMDFG.CDG
NI211125
SI TECHNICAL ERROR
```

### Delayed Takeoff
#### Generic
```
MVT
SD200/21.PMDFG.CDG
AD1115 EO1135
SI DEICING
```
#### With ETA
```
MVT
SD200/21.PMDFG.CDG
AD1115 EO1135 EA1530 FRA
DL72/0025
```

### Return to Ramp
```
MVT
SD200/21.PMDFG.CDG
AD1115 RR1125
DL72/0015
SI ABORTED TAKEOFF
```
### Return from Airborne
```
MVT
SD200/21.PMDFG.CDG
FR1200/1215
SI BIRDSTRIKE
```
### Revised ETA
Used while airborne to update destination station if changes from departure message.
```
MVT
SD200/21.PMDFG.FRA
EA1515
EB1525
SI AVOIDING WEATHER
```
### Arrival taxi time variance
Used after landing, but if the taxi time from the RWY to the parking deviates from original/normal plan.
```
MVT
SD200/21.PMDFG.FRA
AA1510 EB1520
SI LONG TAXI
```

### Additional
All messages above can be corrected when starting with `COR`. This message normally overwrites the original message. 
```
COR
MVT
SD200/21.PMDFG.CDG
AD1059/1115 EA1500 FRA
DL72/0014
PX112
SI CORRECTION TO HAVE SHORTER DELAY ;) 
```

One could also send the same message twice, this would result in a `PDM`- Possible Duplicate Message. 
```
PDM
MVT
SD200/21.PMDFG.CDG
AD1059/1115 EA1500 FRA
DL72/0014
PX112
SI CORRECTION TO HAVE SHORTER DELAY ;) 
```