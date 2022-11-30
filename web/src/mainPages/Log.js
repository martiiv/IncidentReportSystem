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

/**
 * Log screen to get information about a specific Incident.
 * @returns {JSX.Element}
 * @constructor
 */
function Log() {
    const [incident, setIncident] = useState([])
    const [lessonsLearned, setLessonsLearned] = useState([])
    const [description, setDescription] = useState("")
    const [receiver, setReceiver] = useState("")
    const [changed, setChanged] = useState(false)
    const [backup, setBackup] = useState("")
    const [newLessonsLearned, setNewLessonsLearned] = useState("")
    const id = getProjectID()
    const navigate = useNavigate();

    /**
     * Function that will split the lessons learned from given format.
     * @param input lessons learned.
     */
    const splitLessonsLearned = (input) => {
        if(input.length > 0){
            let array = input.split(';')
            setLessonsLearned(array)
        }
    }

    /**
     * Function that will fetch given incident.
     */
    useEffect(() => {
        const fetch = async () => {
            await fetchData(INCIDENT_URL + "?id=" + id).then(data => {
                setIncident(data.data)
                splitLessonsLearned(data.data.lessonlearned)
                setBackup(data.data.lessonlearned)
                setDescription(data.data.description)
                setReceiver(data.data.receivingGroup)
            })

        }

        fetch()
            // make sure to catch any error
            .catch(() => alert("No connection, try again later"));
    }, [])


    /**
     * Function that will update the lessons learned of the incident.
     * @returns {Promise<void>} result of the promise function.
     */
    const handleChange = async () => {
        const lessonLearnedString = lessonsLearned.join(';')
        const body = {
            incidentId: parseInt(incident.id),
            lessonlearned: lessonLearnedString

        }
        await putModel(INCIDENT_URL, body)
            .then(() => {
                navigate(-1)
            })
            .catch(res => console.log(res))

    }

    /**
     * Function that will add the new lessons learned.
     */
    const addLesson = () => {
        setChanged(true)
        setLessonsLearned(prevState => ([...prevState, newLessonsLearned]))
        setNewLessonsLearned("")
    }

    /**
     * Function that will delete a specific lessons learned.
     *
     * @param e the specific lesson
     */
    const deleteLesson = (e) => {
        // eslint-disable-next-line no-restricted-globals
        if (confirm("Do you want to delete item?")) {
            setLessonsLearned(lessonsLearned.filter(item => item !== e))
            setChanged(true)
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
                            <textarea placeholder={"Enter lesson learned"}
                                      className={"textarea-log"}
                                      onChange={(e) => {
                                          setNewLessonsLearned(e.target.value)
                                      }}
                                      value={newLessonsLearned}
                                      onKeyDown={event => {
                                          if (event.key === 'Enter' && !event.shiftKey) {
                                              event.preventDefault()
                                              addLesson()
                                          }
                                      }}
                            />
                            <button className={"btn send-btn"} onClick={addLesson}>Add</button>
                        </div>
                    </div>

                    <div className={"wrapper countermeasures"}>
                        <ul className={"countermeasure-list"}>
                            {lessonsLearned.length > 0 ? lessonsLearned.map(item =>
                                <div className={"countermeasure-list-outer"}>
                                    <li className={"countermeasure-list-element"}>
                                        {item}
                                    </li>
                                    <button className={"btn delete-button img"} onClick={() => deleteLesson(item)}>
                                        <img src={trash_can} alt={"trash_icon"}/>
                                        Delete
                                    </button>
                                </div>
                            ) : <div className={"countermeasure-list-outer"}>
                                    No "lessons learned" in system
                            </div>}
                        </ul>
                    </div>
                </div>
            </div>


            <div className={"log-btn"}>
                <button className={"btn cancel-btn"} onClick={() => setLessonsLearned(backup)}>Cancel</button>
                {changed ? <button onClick={handleChange} className={"btn"}>Save</button> :
                    <button className={"disabled-btn"}>Save</button>}

            </div>

        </div>
    );
}

export default Log;
