import React from "react"
import DummyData from "../constants/DummyData";


/**
 * Function that will get the project id form the url.
 * @returns {string} project id
 */
function getProjectID() {
    const pathSplit = window.location.href.split("/")
    return pathSplit[pathSplit.length - 1]
}

function getDataFromID(id){
    const data = DummyData;
    console.log(data)
    return (data.filter( incident => String(incident.id) === id))
}

function Log() {
    const id = getProjectID()
    const incident = getDataFromID(id)[0]
    return (
        <div>
            <h1 className={"incident_title"}>{incident.name} </h1>
            <div className={"date-sender"} style={{display: "flex", gap: "20px"}}>
                <h3 style={{}}>Date: {incident.date}</h3>
                <h3>Sender: {incident.systemManager}</h3>
            </div>

        </div>
    );

}

export default Log;
