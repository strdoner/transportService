// src/App.js
import React from 'react';
import Vehicles from './components/Vehicles';
import ParkingList from './components/ParkingList';
import './App.css';

function App() {
    return (
        <div className="App">
            <header className="header">
                <h1>Умный город</h1>
                <p>Транспорт и парковки</p>
            </header>

            <div className="grid-container">
                <div className="section">
                    <Vehicles />
                </div>
                <div className="section">
                    <ParkingList />
                </div>
            </div>

            <footer className="footer">
                Transport Service · {new Date().getFullYear()}
            </footer>
        </div>
    );
}

export default App;