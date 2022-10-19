import React from "react";
import {
    AppBar, Toolbar, Button
} from '@material-ui/core';
import {Link} from "react-router-dom";
import {DASHBOARD_URL, INCIDENT_URL, MANAGER_URL, RECEIVER_URL} from "../constants/WebURL";


/**
 Component that will be used as a top bar for the user to navigate throughout the application.
 */

const TopBar = () => {
        //Top bar with interactive buttons to navigate.
        return (
            <AppBar position="sticky">
                <Toolbar className="toolbar">
                    <Link className="link" to={DASHBOARD_URL}>
                        <Button className="topBarBtn">Dashboard</Button>
                    </Link>
                    <Link className="link" to={INCIDENT_URL}>
                        <Button className="topBarBtn">Incident</Button>
                    </Link>
                    <Link className="link" to={RECEIVER_URL}>
                        <Button className="topBarBtn">Warning Receiver</Button>
                    </Link>
                    <Link className="link" to={MANAGER_URL}>
                        <Button className="topBarBtn">Profile</Button>
                    </Link>
                </Toolbar>
            </AppBar>
        );
}

export default TopBar;
