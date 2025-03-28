const picker = document.querySelector('gmpx-place-picker');
const result = document.querySelector('.result');
picker.addEventListener('gmpx-placechange', (e) => {
    console.log(e.detail);
    result.textContent = e.target.value?.formattedAddress ?? '';
});
