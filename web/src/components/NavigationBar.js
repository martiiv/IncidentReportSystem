import React from "react";
import {
    AppBar, Toolbar, Button
} from '@material-ui/core';
import {Link} from "react-router-dom";
import {DASHBOARD_URL, INCIDENT_URL, MANAGER_URL, RECEIVER_URL} from "../constants/WebURL";
import "../index.css"

/**
 */

/**
 * Component that will be used as a top bar for the user to navigate throughout the application.
 * The component uses Appbar from MaterialUI
 *
 * @returns {JSX.Element} the Top bar
 */
const TopBar = () => {
        return (
            <AppBar position="sticky" className={"mainAppBar"}>
                <Toolbar className="toolbar">
                    <Link className="link" to={DASHBOARD_URL}>
                        <Button className="topBarBtn">Dashboard</Button>
                    </Link>
                    <Link className="link"  to={INCIDENT_URL}>
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
