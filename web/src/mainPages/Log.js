import React, {useEffect, useState} from "react"
import {useNavigate} from 'react-router-dom';
import "./Log.css"
import fetchData from "../middleware/FetchData";
import {INCIDENT_URL} from "../constants/WebURL";
import putModel from "../middleware/putData";
import "./Create.css"
import calender from "../images/calendar.png"
import sender from "../images/postman.png"
import tag from "../images/price-tag.png"
import group from "../images/group.png"


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
    const [, setTags] = useState([])
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
                    <div className={"card card-small"}>
                        <img src={calender} className={"image-style"} alt={""}/>
                        <label className={"beside-image"}>{incident.date}</label>
                    </div>

                    <div className={"card card-small"}>
                        <img src={sender} className={"image-style"} alt={""}/>
                        <label className={"beside-image"}>{incident.sendbymanager}</label>
                    </div>

                    <div className={"card card-small"}>
                        <img src={tag} className={"image-style"} alt={""}/>
                        <label className={"beside-image"}>{incident.tag}</label>

                    </div>
                    <div className={"card card-small"}>
                        <img src={group} className={"image-style"} alt={""}/>
                        <label className={"beside-image"}>{receiver}</label>
                    </div>
                </div>

                <div className={"card"}>
                    <h2 className={"text-area-title"}>Description</h2>
                    <p>{description}</p>
                </div>


                <div className={"card"}>
                    <h2 className={"text-area-title"}>Countermeasures</h2>
                    <div className={"flex-div"}>
                        <ul style={{width: "50%"}}>
                            {countermeasure.map(item =>
                                <li>
                                    {item}
                                    <button onClick={() => deleteCountermeasure(item)}>Delete</button>
                                </li>
                            )}
                        </ul>



                        <div className={"text-and-btn"}>
                            <div style={{display: "flex", flexDirection: "column"}}>
                        <textarea placeholder={"Enter new countermeasure"}
                                  style={{ resize: "none", width: "600px" , minHeight: "100px", maxHeight: "100px"}}
                                  className={"textarea-log"}
                                  onChange={(e) => {
                                      setNewCountermeasure(e.target.value)
                                  }}
                                  value={newCountermeasure}
                                  onKeyDown={event => {
                                      if (event.key === 'Enter') {
                                          addCountermeasure()
                                      }
                                  }}
                        />
                            <button className={"btn send-btn"} onClick={addCountermeasure}>Add</button>
                        </div>
                        </div>
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
