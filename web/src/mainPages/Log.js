import React from "react"
import DummyData from "../constants/DummyData";
import { useNavigate } from 'react-router-dom';
import "./Log.css"

/**
 * Function that will get the project id form the url.
 * @returns {string} project id
 */
function getProjectID() {
    const pathSplit = window.location.href.split("/")
    return pathSplit[pathSplit.length - 1]
}

function getDataFromID(id) {
    const data = DummyData;
    console.log(data)
    return (data.filter(incident => String(incident.id) === id))
}

function Log() {
    const id = getProjectID()
    const incident = getDataFromID(id)[0]
    const navigate = useNavigate();
    return (
        <div className={"log-screen"}>
            <button className={"btn"} onClick={() => navigate(-1)}>Back</button>
            <h1 className={"incident_title"}>{incident.name}</h1>
            <div className={"date-sender"}>
                <h3>Date: {incident.date}</h3>
                <h3>Sender: {incident.systemManager}</h3>
            </div>
            <div className={"text-area-wr"}>
                {/*<div>
                    <h4>Warning Receiver</h4>
                    <textarea>
                        {incident.receivingGroup}
                    </textarea>
                </div>*/}
                <div>
                    <h4 className={"text-area-title"}>Description</h4>
                    <textarea className={"textarea-log textarea-context"}>
                        {incident.context}
                    </textarea>
                </div>
            </div>
            <div className={"text-area-wr"} >
                <div>
                    <h4 className={"text-area-title"}>Countermeasures</h4>
                    <textarea className={"textarea-log textarea-countermeasure"}>
                        {incident.countermeasure}
                    </textarea>
                </div>
            </div>
            <div className={"log-btn"}>
                <button className={"btn cancel-btn"}>Cancel</button>
                <button className={"btn"}>Save</button>

            </div>
        </div>
    );
}

export default Log;
