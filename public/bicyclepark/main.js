/**
 * @license
 * Copyright 2024 Google LLC. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 */
let map;
let infoWindow;

async function init() {
  const {InfoWindow} = await google.maps.importLibrary("maps");

  map = document.querySelector('gmp-map').innerMap;
  infoWindow = new InfoWindow({pixelOffset: {height: -37}});

  // Get the earthquake data (JSONP format).
  // This feed is a copy from the USGS feed, you can find the originals here:
  //   http://earthquake.usgs.gov/earthquakes/feed/v1.0/geojson.php
  const script = document.createElement("script");
  // script.src = "https://storage.googleapis.com/mapsdevsite/json/quakes.geo.json";
  script.src="bicyclepark/ltaBicycleRack.geojson"
  document.head.appendChild(script);
}

function showRackInfo(position, feature) {
  const content = `
    <div style="padding: 8px">
      <h2 style="margin-top: 0">${feature.getProperty('Description')}</h2>
    </div>
  `;

  infoWindow.setOptions({content, position});
  infoWindow.open({map, shouldFocus: false});
}

// Defines the callback function referenced in the jsonp file.
window.eqfeed_callback = (data) => {
  map.data.addGeoJson(data);
  map.data.setStyle((feature) => ({
    title: feature.getProperty('place')
  }));
  map.data.addListener('click', (e) => showRackInfo(e.latLng, e.feature));
}

init();