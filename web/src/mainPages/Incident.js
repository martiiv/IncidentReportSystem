import React, {useEffect, useState} from "react";
import postModel from "../middleware/postData";
import {DASHBOARD_URL, GROUPS_URL, INCIDENT_URL, TAG_Query} from "../constants/WebURL";
import TagsInput from "../components/TagsInput";
import GroupSelectComponent from "../components/GroupSelectComponent";
import "./Create.css";
import fetchData from "../middleware/FetchData";
import {ReactNotifications} from 'react-notifications-component'
import 'react-notifications-component/dist/theme.css'
import {Store} from 'react-notifications-component';
import {useNavigate} from "react-router-dom";


function Incident() {
    const credentials = (JSON.parse(sessionStorage.getItem("credentials")))

    const [stateTest, setStateTest] = useState({
        name: "",
        tag: "",
        description: "",
        company: credentials.company,
        receivingGroup: "",
        countermeasure: "",
        sendbymanager: credentials.userName
    })

    const [tags, setTags] = useState("")
    const [selectedGroups, setSelectedGroups] = useState(0)
    const [isPending, setIsPending] = useState(false);
    const [tagsOption, setTagsOption] = useState([])
    const [groupsOption, setgroupsOption] = useState([])
    const navigate = useNavigate();


    useEffect(() => {
        const fetchTags = async () => {
            let counter = 0

            await fetchData(INCIDENT_URL + TAG_Query)
                .then(res =>
                    (setTagsOption(res.data.map(item => ({id: counter++, name: item.tag})))))
                .catch(e => console.log(e))
        }

        const fetchGroups = async () => {
            let counter = 0
            await fetchData(GROUPS_URL)
                .then(res =>
                    (setgroupsOption(res.data.map(item => ({id: counter++, name: item.name})))))
                .catch(e => console.log(e))
        }

        fetchTags().catch(e => console.log(e))
        fetchGroups().catch(e => console.log(e))
    }, [])


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

    console.log(stateTest)

    function handleChange(evt) {
        const value = evt.target.value;
        setStateTest({
            ...stateTest, [evt.target.name]: value
        })

        setStateTest(prevState => ({
            ...prevState,
            "tag": tags,
            "receivingGroup": selectedGroups,
        }))

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
                    {tagsOption.length > 0 ? <TagsInput
                        setTagsFunc={setTags}
                        data={tagsOption}
                    /> : null}
                </label>

                {groupsOption.length > 0 ?
                    <GroupSelectComponent
                        data={groupsOption}
                        setSelectedFunc={setSelectedGroups}

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
                    <input
                        type={"text"}
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
