import React, {useEffect, useState} from "react"
import DummyData from "../constants/DummyData";
import { useNavigate } from 'react-router-dom';
import "./Log.css"
import TagsInput from "../components/TagsInput";
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

function getDataFromID(id) {
    const data = DummyData;
    console.log(data)
    return (data.filter(incident => String(incident.id) === id))
}



function Log() {
    const [incident, setIncident] = useState([])
    const [tags, setTags] = useState([])
    const [countermeasure, setCountermeasure] = useState("")
    const [description, setDescription] = useState("")
    const [receiver, setReceiver] = useState("")

    const id = getProjectID()


    useEffect( () => {
        const fetch = async () => {
            const data = await fetchData(INCIDENT_URL + "?id=" + id)
            setIncident(data.data)
            setCountermeasure(data.data.countermeasure)
            setDescription(data.data.description)
            setReceiver(data.data.receivingGroup)
            setTags(data.data.tags)
        }

        fetch()
            // make sure to catch any error
            .catch(console.error);
    }, [])


    console.log(incident);

    const navigate = useNavigate();


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
                    <textarea className={"textarea-small"} value={receiver}>
                    </textarea>
                </div>
                <div>
                    <h4 className={"text-area-title"}>Description</h4>
                    <textarea className={"textarea-log textarea-context"} value={description} onChange={e => setDescription(e.target.value)}>
                    </textarea>
                </div>
            </div>
            <div className={"text-area-wr"}>
                <div>
                    <h4 className={"text-area-title tag-title"}>Tags</h4>
                    <div className={"tag-input textarea-small"}>
                        {incident.tag ?  <TagsInput setTagsFunc={setTags}
                                           data={incident.tag}/> : null}

                    </div>
                </div>

                <div>
                    <h4 className={"text-area-title"}>Countermeasures</h4>
                    <textarea className={"textarea-log textarea-countermeasure"} value={countermeasure} onChange={e => setCountermeasure(e.target.value)}>
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
