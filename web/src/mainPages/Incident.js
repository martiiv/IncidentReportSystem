import React, {useEffect, useState} from "react";
import postModel from "../middleware/postData";
import {COUNTERMEASURE, DASHBOARD_URL, GROUPS_URL, INCIDENT_URL} from "../constants/WebURL";
import TagsInput from "../components/TagsInput";
import GroupSelectComponent from "../components/GroupSelectComponent";
import "./Create.css";
import fetchData from "../middleware/FetchData";
import {ReactNotifications} from 'react-notifications-component'
import 'react-notifications-component/dist/theme.css'
import {Store} from 'react-notifications-component';
import {useNavigate} from "react-router-dom";

/**
 * Incident page, to create and send new incident.
 * @returns {JSX.Element}
 * @constructor
 */
function Incident() {
    const credentials = (JSON.parse(sessionStorage.getItem("credentials")))

    const [stateTest, setStateTest] = useState({
        name: "",
        tag: "",
        description: "",
        company: credentials.company,
        countermeasure: "",
        receivingGroup: "",
        sendbymanager: credentials.userName
    })

    const [setTags] = useState("")
    const [isPending, setIsPending] = useState(false);
    const [tagsOption, setTagsOption] = useState([])
    const [groupsOption, setgroupsOption] = useState([])

    const navigate = useNavigate();


    /**
     * Function to fetch countermeasures, and groups, from database.
     */
    useEffect(() => {
        const fetchTags = async () => {
            let counter = 0
            await fetchData(COUNTERMEASURE)
                .then(res => {
                    (setTagsOption(res.data.map(item => ({id: counter++, name: item.tag, description: item.description}))))

                })
        }

        const fetchGroups = async () => {
            let counter = 0
            await fetchData(GROUPS_URL)
                .then(res =>
                    (setgroupsOption(res.data.map(item => ({id: counter++, name: item.name})))))
        }

        fetchTags().catch(() => alert("No connection, try again later"))
        fetchGroups().catch(() => alert("No connection, try again later"))
    }, [])


    /**
     * Function that will submit the Incident to the database.
     * Will Send confirmation notification on success and rejection.
     *
     * @param e data sent from forms
     * @returns {Promise<void>}
     */
    const handleSubmit = async (e) => {
        e.preventDefault()
        setIsPending(true)
        await postModel(INCIDENT_URL, JSON.stringify(stateTest))
            .then(() => {
                setIsPending(false)
                Store.addNotification({
                    title: "Successfully sent!",
                    message: "Incident is reported to the users",
                    type: "success",
                    insert: "top",
                    container: "top-right",
                    animationIn: ["animate__animated", "animate__fadeIn"],
                    animationOut: ["animate__animated", "animate__fadeOut"],
                    dismiss: {
                        duration: 2000,
                        onScreen: true
                    }
                });
            })
            .then(setTimeout(() => {
                navigate(DASHBOARD_URL)
            }, 2000))
            .catch(() => {
                setIsPending(false)
                alert("Incident not sent. please try again")
            })
    }

    /**
     * Function that will set the given value to the Incident.
     *
     * @param evt input from input field.
     */
    function handleChange(evt) {
        const value = evt.target.value;
        setStateTest({
            ...stateTest, [evt.target.name]: value
        })
    }

    return (
        <div className={"create"}>
            <form className={"create-forms"} onSubmit={handleSubmit} onKeyDown={(e) => {
                e.key === 'Enter' && e.preventDefault();
            }}>
                <h2>New Incident</h2>
                <ReactNotifications/>
                <label>Incident Title:
                    <input
                        type={"text"}
                        name={"name"}
                        required
                        value={stateTest.name}
                        onChange={handleChange}
                    />
                </label>

                <label>Tags:
                    {tagsOption.length > 0 ?
                        <TagsInput
                        setTagsFunc={setTags}
                        data={tagsOption}
                        setStateTestFunc={setStateTest}
                        /> : null}
                </label>

                {groupsOption.length > 0 ?
                    <GroupSelectComponent
                        data={groupsOption}
                        setStateTestFunc={setStateTest}

                    /> : null
                }

                <label>Incident description:
                    <textarea
                        name={"description"}
                        required
                        value={stateTest.description}
                        onChange={handleChange}
                    />
                </label>


                <label>Countermeasures
                    <textarea
                        name={"countermeasure"}
                        value={stateTest.countermeasure}
                        onChange={handleChange}
                    />
                </label>
                <div style={{display: "flex", gap: "50px"}}>
                    <button className={"btn"} onClick={() => navigate(-1)}>Back</button>
                    {!isPending && <button className={"btn send-btn"}>SEND</button>}
                    {isPending && <button className={"btn send-btn"} disabled>Sending mail...</button>}
                </div>
            </form>
        </div>
    )
}

export default Incident;
