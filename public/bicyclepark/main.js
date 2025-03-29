/**
 * @license
 * Copyright 2024 Google LLC. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 */
let map;
let infoWindow;
let whitePin;

async function loadData() {
  const { InfoWindow } = await google.maps.importLibrary("maps");
  const { PinElement } = await google.maps.importLibrary("marker")
  whitePin = new PinElement({ glyphColor: "white", background: "#6600cc" });

  map = document.querySelector('gmp-map').innerMap;
  await setInitialMapPosition()
  infoWindow = new InfoWindow({ pixelOffset: { height: -37 } });
  const ctr = map.getCenter()
  map.data.loadGeoJson(`bicycleParkingSpots?Lat=${ctr.lat()}&Lng=${ctr.lng()}`)

  map.data.addListener('click', (e) => showRackInfo(e.latLng, e.feature));
}

function showRackInfo(position, feature) {
  const content = `
    <div>
      <strong>${feature.getProperty('Description')}</strong><br/>
      <span>${feature.getProperty("RackType")}</span><br/>
      <span>${feature.getProperty("RackCount")} Lots</span><br/>
      <span>Sheltered: ${feature.getProperty("ShelterIndicator")}</span><br/>
    </div>
  `;

  infoWindow.setOptions({ content, position });
  infoWindow.open({ map, shouldFocus: false });
}

async function init() {
  await loadData()

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
    const loc = place.location
    map.innerMap.data.loadGeoJson(`bicycleParkingSpots?Lat=${loc.lat()}&Lng=${loc.lng()}`)

    marker.position = place.location;
    marker.content = whitePin.element;
    infowindow.setContent(
      `<strong>${place.displayName}</strong><br>
       <span>${place.formattedAddress}</span><br>
       <span>${place.location}</span>
    `);
    infowindow.open(map.innerMap, marker);
  });

  marker.addEventListener("gmp-click", () => {
    const place = placePicker.value
    if (!place) { return }

    marker.content = whitePin.element
    infowindow.setContent(
      `<strong>${place.displayName}</strong><br>
       <span>${place.formattedAddress}</span><br>
       <span>${place.location}</span>
    `)
    infowindow.open(map.innerMap, marker)
  })
}

document.addEventListener('DOMContentLoaded', init);

async function setInitialMapPosition() {
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(position => {
      const marker = document.querySelector('gmp-advanced-marker');
      const pos = {
        lat: position.coords.latitude,
        lng: position.coords.longitude,
      }
      map.setCenter(pos)
      marker.position = pos
      marker.content = whitePin.element
    })
  }
}