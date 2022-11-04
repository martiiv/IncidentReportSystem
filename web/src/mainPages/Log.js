import React, {useEffect, useState} from "react"
import DummyData from "../constants/DummyData";
import { useNavigate } from 'react-router-dom';
import "./Log.css"
import TagsInput from "../components/TagsInput";
import fetchData from "../middleware/FetchData";
import {INCIDENT_URL} from "../constants/WebURL";
import putModel from "../middleware/putData";

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
    const [tags, setTags] = useState([])
    const [countermeasure, setCountermeasure] = useState("")
    const [description, setDescription] = useState("")
    const [receiver, setReceiver] = useState("")
    const [changed, setChanged] = useState(false)
    const [backup, setBackup] = useState("")


    const id = getProjectID()


    useEffect( () => {
        const fetch = async () => {
            const data = await fetchData(INCIDENT_URL + "?id=" + id)
            setIncident(data.data)
            setCountermeasure(data.data.countermeasure)
            setBackup(data.data.countermeasure)

            setDescription(data.data.description)
            setReceiver(data.data.receivingGroup)
            setTags(data.data.tags)
        }

        fetch()
            // make sure to catch any error
            .catch(console.error);
    }, [])


    const navigate = useNavigate();

    const handleChange = async () => {

        const body = {
            incidentId: parseInt(incident.id),
            countermeasure: countermeasure
        }
        console.log(JSON.stringify(body))
        await putModel(INCIDENT_URL, body)
            .then(res => console.log(res.statusCode))
            .catch(res => console.log(res))

    }

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
                    <textarea className={"textarea-log textarea-context"} value={description}>
                    </textarea>
                </div>
            </div>
            <div className={"text-area-wr"}>
                <div>
                    <h4 className={"text-area-title tag-title"}>Tags</h4>
                    <div className={"tag-input textarea-small"}>
                        {incident.tag ?  <TagsInput setTagsFunc={setTags}
                                                    data={incident.tag}
                                                    changeable={false}
                        /> : null}

                    </div>
                </div>

                <div>
                    <h4 className={"text-area-title"}>Countermeasures</h4>
                    <textarea className={"textarea-log textarea-countermeasure"} value={countermeasure} onChange={e => {
                        setChanged(true)
                        setCountermeasure(e.target.value)
                    }}>
                    </textarea>
                </div>

            </div>
            <div className={"log-btn"}>
                <button className={"btn cancel-btn"} onClick={() => setCountermeasure(backup)}>Cancel</button>
                {changed ? <button onClick={handleChange} className={"btn"}>Save</button> : <button  className={"disabled-btn"}>Save</button> }

            </div>



        </div>
    );
}

export default Log;
