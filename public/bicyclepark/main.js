/**
 * @license
 * Copyright 2024 Google LLC. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 */
let map;
let infoWindow;

async function init() {
  const { InfoWindow } = await google.maps.importLibrary("maps");

  map = document.querySelector('gmp-map').innerMap;
  infoWindow = new InfoWindow({ pixelOffset: { height: -37 } });

  // Get the earthquake data (JSONP format).
  // This feed is a copy from the USGS feed, you can find the originals here:
  //   http://earthquake.usgs.gov/earthquakes/feed/v1.0/geojson.php
  const script = document.createElement("script");
  // script.src = "https://storage.googleapis.com/mapsdevsite/json/quakes.geo.json";
  script.src = "bicyclepark/ltaBicycleRack.geojson"
  document.head.appendChild(script);
}

function showRackInfo(position, feature) {
  const content = `
    <div style="padding: 8px">
      <h2 style="margin-top: 0">${feature.getProperty('Description')}</h2>
    </div>
  `;

  infoWindow.setOptions({ content, position });
  infoWindow.open({ map, shouldFocus: false });
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

///////////////////////////////////////////////////////////////////

async function init2() {
  console.log("init2")
  await customElements.whenDefined('gmp-map');

  const map = document.querySelector('gmp-map');
  const marker = document.querySelector('gmp-advanced-marker');
  const placePicker = document.querySelector('gmpx-place-picker');
  const infowindow = new google.maps.InfoWindow();

  map.innerMap.setOptions({
    mapTypeControl: false
  });

  placePicker.addEventListener('gmpx-placechange', () => {
    const place = placePicker.value;

    if (!place.location) {
      window.alert(
        "No details available for input: '" + place.name + "'"
      );
      infowindow.close();
      marker.position = null;
      return;
    }

    if (place.viewport) {
      map.innerMap.fitBounds(place.viewport);
    } else {
      map.center = place.location;
      map.zoom = 17;
    }

    marker.position = place.location;
    infowindow.setContent(
      `<strong>${place.displayName}</strong><br>
       <span>${place.formattedAddress}</span>
    `);
    infowindow.open(map.innerMap, marker);
  });
}

document.addEventListener('DOMContentLoaded', init2);