import React, {useEffect, useState} from "react"
import DummyData from "../constants/DummyData";
import { useNavigate } from 'react-router-dom';
import "./Log.css"
import TagsInput from "../components/TagsInput";
import dummyData from "../constants/DummyData";
import fetchData from "../middleware/FetchData";
import {INCIDENT_URL} from "../constants/WebURL";

/**
 * Function that will get the project id form the url.
 * @returns {string} project id
 */
function getProjectID() {
    const pathSplit = window.location.href.split("/")
    return pathSplit[pathSplit.length - 1]
}



function Log() {
    const [incident, setIncident] = useState([])
    const id = getProjectID()


    useEffect( () => {
        const fetch = async () => {
            const data = await fetchData(INCIDENT_URL + "?id=" + id)
            setIncident(data.data)
        }

        fetch()
            // make sure to catch any error
            .catch(console.error);
    }, [])


    console.log(incident);

    const navigate = useNavigate();
    const [tags, setTags] = useState([])


    return (
        <div className={"log-screen"}>
            <button className={"btn"} onClick={() => navigate(-1)}>Back</button>
            <h1 className={"incident_title"}>{incident.name}</h1>
            <div className={"date-sender"}>
                <h3>Date: {incident.date}</h3>
                <h3>Sender: {incident.sendbymanager}</h3>
            </div>
            <div className={"text-area-wr"}>
                <div className={"warning-receiver"}>
                    <h4 className={"text-area-title"}>Warning Receiver</h4>
                    <textarea className={"textarea-small"}>
                        {incident.receivingGroup}
                    </textarea>
                </div>
                <div>
                    <h4 className={"text-area-title"}>Description</h4>
                    <textarea className={"textarea-log textarea-context"}>
                        {incident.description}
                    </textarea>
                </div>
            </div>
            <div className={"text-area-wr"}>

                <div>
                    <h4 className={"text-area-title tag-title"}>Tags</h4>
                    <div className={"tag-input textarea-small"}>
                        <TagsInput setTagsFunc={setTags}
                                   data={incident.tags}/>
                    </div>
                </div>

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
