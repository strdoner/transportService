import React, {useState} from 'react';
import Vehicles from './components/Vehicles';
import ParkingList from './components/ParkingList';
import ReserveForm from "./components/ReserveForm";
import './styles/style.css'

function App() {
    const [vehicles, setVehicles] = useState([{"id":1,"name":"Toyota Camry","type":"car","latitude":55.749823,"longitude":37.618942,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":2,"name":"Volkswagen Tiguan","type":"car","latitude":55.752104,"longitude":37.621567,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":3,"name":"Hyundai Solaris","type":"car","latitude":55.756321,"longitude":37.61423,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":4,"name":"Kia Rio","type":"car","latitude":55.748765,"longitude":37.62541,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":5,"name":"KAMAZ 6520","type":"truck","latitude":55.732145,"longitude":37.64231,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":6,"name":"MAN TGS","type":"truck","latitude":55.768901,"longitude":37.598765,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":7,"name":"Yamaha MT-07","type":"motorcycle","latitude":55.751234,"longitude":37.619876,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":8,"name":"Harley-Davidson Street 750","type":"motorcycle","latitude":55.754321,"longitude":37.622109,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":9,"name":"BMW R 1250 GS","type":"motorcycle","latitude":55.747654,"longitude":37.616543,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":10,"name":"Stels Navigator 500","type":"bike","latitude":55.753421,"longitude":37.620123,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":11,"name":"Merida Big.Seven 40","type":"bike","latitude":55.758765,"longitude":37.612345,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":12,"name":"Giant Escape 3","type":"bike","latitude":55.745678,"longitude":37.628765,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":13,"name":"Xiaomi Mi Electric Scooter Pro 2","type":"scooter","latitude":55.750987,"longitude":37.617654,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":14,"name":"Segway Ninebot MAX","type":"scooter","latitude":55.755432,"longitude":37.623456,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":15,"name":"Inokim Light 2","type":"scooter","latitude":55.749012,"longitude":37.626789,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":16,"name":"Lada Granta","type":"car","latitude":55.760123,"longitude":37.609876,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":17,"name":"Volvo FH16","type":"truck","latitude":55.725678,"longitude":37.651234,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":18,"name":"Ducati Monster","type":"motorcycle","latitude":55.75789,"longitude":37.618765,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":19,"name":"Cannondale Trail 5","type":"bike","latitude":55.752345,"longitude":37.627654,"updated_at":{"seconds":1764394324,"nanos":834889000}},{"id":20,"name":"NIU NQi GT","type":"scooter","latitude":55.756789,"longitude":37.615432,"updated_at":{"seconds":1764394324,"nanos":834889000}}]);
    const [parkings, setParkings] = useState([{"id":1,"name":"Парковка у Кремля","capacity":12,"latitude":55.75222,"longitude":37.61556,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":2,"name":"Государственная парковка на Тверской","capacity":20,"latitude":55.76234,"longitude":37.60987,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":3,"name":"Паркинг у метро «Охотный Ряд»","capacity":18,"latitude":55.75712,"longitude":37.61765,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":4,"name":"Бизнес-парковка «Славянка»","capacity":25,"latitude":55.73987,"longitude":37.62891,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":5,"name":"Парковка у ГУМа","capacity":8,"latitude":55.75432,"longitude":37.6231,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":6,"name":"ТЦ «Атриум» — подземная парковка","capacity":30,"latitude":55.76345,"longitude":37.6421,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":7,"name":"Парковка на Красной Пресне","capacity":15,"latitude":55.75123,"longitude":37.5689,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":8,"name":"Паркинг у метро «Киевская»","capacity":22,"latitude":55.74678,"longitude":37.56321,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":9,"name":"Офисный паркинг «Москва-Сити»","capacity":28,"latitude":55.7489,"longitude":37.53456,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":10,"name":"Парковка у ВДНХ","capacity":20,"latitude":55.82345,"longitude":37.6389,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":11,"name":"Парковка на Павелецкой","capacity":10,"latitude":55.73123,"longitude":37.63567,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":12,"name":"Гостевая парковка «Центральная»","capacity":14,"latitude":55.75567,"longitude":37.62123,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":13,"name":"Паркинг у Лужников","capacity":16,"latitude":55.73456,"longitude":37.58123,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":14,"name":"Парковка у Садового кольца","capacity":12,"latitude":55.76789,"longitude":37.61234,"created_at":{"seconds":1764394324,"nanos":839435000}},{"id":15,"name":"Мини-парковка на Арбате","capacity":6,"latitude":55.74987,"longitude":37.58765,"created_at":{"seconds":1764394324,"nanos":839435000}}])

    const [selectedVehicleId, setSelectedVehicleId] = useState(null);
    const [selectedParkingId, setSelectedParkingId] = useState(null);

    return (
        <div className="App">
            <header className="header">
                <h1 className="text-center m-1">Умный город</h1>
                <h4 className="text-center text-muted me-4">Модуль бронирования парковочных мест</h4>
            </header>
            <div className="container">
                <div className="row justify-content-center gap-2">
                    <div className="col-5 card">
                        <div className="card-header bg-white pt-2"><h2><i className="bi bi-p-square me-2"></i>Доступные парковки</h2></div>
                        <div className="card-body">
                            <ParkingList parkings={parkings}
                                         setParkings={setParkings}
                                         onParkingSelect={setSelectedParkingId}
                                         selectedId={selectedParkingId}
                            />
                        </div>
                    </div>
                    <div className="col-5 card">
                        <div className="card-header bg-white pt-2"><h2><i className="bi bi-truck me-2"></i>Транспорт</h2></div>
                        <div className="card-body">
                            <Vehicles vehicles={vehicles}
                                      setVehicles={setVehicles}
                                      onVehicleSelect={setSelectedVehicleId}
                                      selectedId={selectedVehicleId}
                            />
                        </div>
                    </div>
                    <div className="col-10 card">
                        <div className="card-header bg-white">
                            <h2>Бронирование парковочного места</h2>
                        </div>
                        <div className="card-body">
                            <ReserveForm vehicles={vehicles}
                                         parkings={parkings}
                                         selectedVehicleId={selectedVehicleId}
                                         selectedParkingId={selectedParkingId}
                            />
                        </div>
                    </div>

                </div>
            </div>


        </div>
    );
}

export default App;