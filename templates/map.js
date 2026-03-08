ymaps.ready(initMap);

function initMap() {
    let myMap = new ymaps.Map("map", {
        center: [55.751574, 37.573856], 
        zoom: 10 
    });

    let carZil131 = new ymaps.Placemark([55.801574, 37.573856], {
        hintContent: 'ЗИЛ-131 (В 011 ВВ 77)',
        balloonContent: '<b>ЗИЛ-131</b><br>Статус: В работе на объекте'
    }, {
        preset: 'islands#blueAutoIcon' 
    });

    let carZil130 = new ymaps.Placemark([55.768778, 37.376940], {
        hintContent: 'ЗИЛ-130 (А 007 АА 77)',
        balloonContent: '<b>ЗИЛ-130</b><br>Статус: На базе'
    }, {
        preset: 'islands#redAutoIcon' 
    });

    let carKamaz = new ymaps.Placemark([55.769320, 37.375968], {
        hintContent: 'КАМАЗ (О 456 ОО 77)',
        balloonContent: '<b>КАМАЗ</b><br>Статус: В ремонте'
    }, {
        preset: 'islands#greenAutoIcon' 
    });

    let carVolvo = new ymaps.Placemark([55.701574, 37.673856], {
        hintContent: 'VOLVO (М 134 ВМ 77)',
        balloonContent: '<b>VOLVO</b><br>Статус: В работе<br>Рейс: Москва - Казань'
    }, {
        preset: 'islands#blueAutoIcon' 
    });

    myMap.geoObjects.add(carZil131);
    myMap.geoObjects.add(carZil130);
    myMap.geoObjects.add(carKamaz);
    myMap.geoObjects.add(carVolvo);
}