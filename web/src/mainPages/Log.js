import React, {useEffect, useState} from "react"
import {useNavigate} from 'react-router-dom';
import "./Log.css"
import fetchData from "../middleware/FetchData";
import {INCIDENT_URL} from "../constants/WebURL";
import putModel from "../middleware/putData";
import "./Create.css"
import trash_can from "../images/trash_can.png";

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
    const [countermeasure, setCountermeasure] = useState([])
    const [description, setDescription] = useState("")
    const [receiver, setReceiver] = useState("")
    const [changed, setChanged] = useState(false)
    const [backup, setBackup] = useState("")
    const [newCountermeasure, setNewCountermeasure] = useState("")


    const id = getProjectID()

    const splitCountermeasure = (input) => {
        let array = input.split(';')
        setCountermeasure(array)
    }

    useEffect(() => {
        const fetch = async () => {
            await fetchData(INCIDENT_URL + "?id=" + id).then(data => {
                setIncident(data.data)
                splitCountermeasure(data.data.countermeasure)
                setBackup(data.data.countermeasure)
                setDescription(data.data.description)
                setReceiver(data.data.receivingGroup)
                setTags(data.data.tags)
            })

        }

        fetch()
            // make sure to catch any error
            .catch(console.error);
    }, [])


    const navigate = useNavigate();
    const handleChange = async () => {
        const countermeasureString = countermeasure.join(';')
        const body = {
            incidentId: parseInt(incident.id),
            countermeasure: countermeasureString
        }
        await putModel(INCIDENT_URL, body)
            .then(() => {
                navigate(-1)
            })
            .catch(res => console.log(res))

    }

    const addCountermeasure = () => {
        setChanged(true)
        setCountermeasure(prevState => ([...prevState, newCountermeasure]))
        setNewCountermeasure("")
    }

    const deleteCountermeasure = (e) => {
        // eslint-disable-next-line no-restricted-globals
        if (confirm("Do you want to delete item?")) {
            setCountermeasure(countermeasure.filter(item => item !== e))
        }
    }


    return (
        <div className={"log-screen"}>
            <button className={"btn"} onClick={() => navigate(-1)}>Back</button>
            <div className={"content"}>
                <h1 className={"incident_title"}>Incident: {incident.name}</h1>
                <div className={"side-by-side"}>
                    <div className={"wrapper calendar"}>
                        <div className={"opacity-layer"}/>
                        <div className={"content-incident calendar"}>{incident.date}</div>
                    </div>

                    <div className="wrapper management">
                        <div className="opacity-layer"/>
                        <div className={"content-incident sender"}>{incident.sendbymanager}</div>
                    </div>

                    <div className={"wrapper incident-tag"}>
                        <div className={"opacity-layer"}/>
                        <div className={"content-incident tag"}>{incident.tag}</div>
                    </div>

                    <div className={"wrapper department"}>
                        <div className={"opacity-layer"}/>
                        <div className={"content-incident department-1"}>{receiver}</div>
                    </div>
                </div>

                <div className={"card-description"}>
                    <div className={"wrapper description-b"}>
                        <div className={"opacity-layer"}/>
                        <div className={"content-incident description-b"}>{description}</div>
                    </div>
                </div>

                <div className={"countermeasures-outer"}>
                    <div className={"wrapper countermeasures"}>
                        <div className={"countermeasure-block"}>
                            <textarea placeholder={"Enter new countermeasure"}
                                      className={"textarea-log"}
                                      onChange={(e) => {
                                          setNewCountermeasure(e.target.value)
                                      }}
                                      value={newCountermeasure}
                                      onKeyDown={event => {
                                          if (event.key === 'Enter' && !event.shiftKey) {
                                              event.preventDefault()
                                              addCountermeasure()
                                          }
                                      }}
                            />
                            <button className={"btn send-btn"} onClick={addCountermeasure}>Add</button>
                        </div>
                    </div>

                    <div className={"wrapper countermeasures"}>
                        <ul className={"countermeasure-list"}>
                            {countermeasure.map(item =>
                                <div className={"countermeasure-list-outer"}>
                                    <li className={"countermeasure-list-element"}>
                                        {item}
                                    </li>
                                    <button className={"btn delete-button img"} onClick={() => deleteCountermeasure(item)}>
                                        <img src={trash_can}/>
                                        Delete
                                    </button>
                                </div>
                            )}
                        </ul>
                    </div>
                </div>
            </div>


            <div className={"log-btn"}>
                <button className={"btn cancel-btn"} onClick={() => setCountermeasure(backup)}>Cancel</button>
                {changed ? <button onClick={handleChange} className={"btn"}>Save</button> :
                    <button className={"disabled-btn"}>Save</button>}

            </div>

        </div>
    );
}

export default Log;
