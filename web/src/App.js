import React, {useState} from "react";
import {Routes, Route} from "react-router-dom";
import Dashboard from "./mainPages/Dashboard";
import {DASHBOARD_URL, ID_PARAM, INCIDENT_URL, MANAGER_URL, RECEIVER_URL} from "./constants/WebURL";
import Incident from "./mainPages/Incident";
import Profile from "./mainPages/Profile";
import WarningReceiver from "./mainPages/WarningReceiver";
import TopBar from "./components/NavigationBar";
import Login from "./mainPages/Login";
import Log from "./mainPages/Log";
import CreateWarningReceiver from "./mainPages/CreateWarningReceiver";
import CreateGroup from "./mainPages/CreateGroup";
import ProtectedRoute from "./mainPages/ProtectedRoute";


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
                    <Route path={DASHBOARD_URL} element={<ProtectedRoute><Dashboard/></ProtectedRoute>}/>
                    <Route path={INCIDENT_URL} element={<ProtectedRoute><Incident/></ProtectedRoute>}/>
                    <Route path={MANAGER_URL} element={<ProtectedRoute><Profile/></ProtectedRoute>}/>
                    <Route path={RECEIVER_URL} element={<ProtectedRoute><WarningReceiver/></ProtectedRoute>}/>
                    <Route path={INCIDENT_URL + ID_PARAM} element={<ProtectedRoute><Log/></ProtectedRoute>}/>
                    <Route path={RECEIVER_URL + "/new"} element={<ProtectedRoute><CreateWarningReceiver/></ProtectedRoute>}/>
                    <Route path={RECEIVER_URL + "/group"} element={<ProtectedRoute><CreateGroup/></ProtectedRoute>}/>
                </Routes>
        </div>
    );
}

export default App;

