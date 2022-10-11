import React from "react";
import {Routes, Route} from "react-router-dom";
import Dashboard from "./mainPages/Dashboard";
import {DASHBOARD_URL, INCIDENT_URL, MANAGER_URL, RECEIVER_URL} from "./constants/WebURL";
import Incident from "./mainPages/Incident";
import Profile from "./mainPages/Profile";
import WarningReceiver from "./mainPages/WarningReceiver";
import TopBar from "./components/NavigationBar";
import Login from "./mainPages/Login";


/**
 * Function that will route the
 * @returns {JSX.Element}
 * @constructor
 */
function App() {
    return (
        <div>
            <TopBar/>
                <Routes> {/*Router that creates the routes the user is able to navigate*/}
                    <Route path={"/"} element={<Login/>}/>
                    <Route path={DASHBOARD_URL} element={<Dashboard/>}/>
                    <Route path={INCIDENT_URL} element={<Incident/>}/>
                    <Route path={MANAGER_URL} element={<Profile/>}/>
                    <Route path={RECEIVER_URL} element={<WarningReceiver/>}/>
                </Routes>
        </div>
    );
}

export default App;

