# GeoJson bridge to Land Transport Autority of Singapore's DataMall

https://datamall.lta.gov.sg/content/datamall/en/dynamic-data.html

## GeoJson feature

```
{
  "type": "Feature",
  "geometry": {
    "type": "Point",
    "coordinates": [103.8514962, 1.2798851] 
  },
  "properties": {
    "name": "Amazon Singapore Office, IOI Central Boulevard."
  }
}
```

## Feature collection

```
{
  "type": "FeatureCollection",
  "features": [{
    "type": "Feature",
    "geometry": {
        "type": "Point",
        "coordinates": [102.0, 0.5]
    },
    "properties": {
        "prop0": "value0"
    }
  }]
}
```
