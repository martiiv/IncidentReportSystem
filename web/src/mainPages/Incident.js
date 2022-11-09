import {useEffect, useState} from "react";
import postModel from "../middleware/postData";
import {GROUPS_URL, INCIDENT_URL, Mail, TAG_Query} from "../constants/WebURL";
import TagsInput from "../components/TagsInput";
import DummyTags from "../constants/DummyTags";
import GroupSelectComponent from "../components/GroupSelectComponent";
import DummyGroups from "../constants/DummyGroups";
import "../components/IncidentReporting.css";
import fetchData from "../middleware/FetchData";

function Incident() {
    const [stateTest, setStateTest] = useState({

        name: "",
        tag: "",
        description: "",
        receivingGroup: "" ,
        countermeasure: "",
    })

    const [tags, setTags] = useState("")
    const [selectedGroups, setSelectedGroups] = useState(0)
    const [isPending, setIsPending] = useState(false);
    const [tagsOption, setTagsOption] = useState([])
    const [groupsOption, setgroupsOption] = useState([])


    useEffect(() => {
        const fetchTags = async () => {
            let counter = 0

            await fetchData(INCIDENT_URL + TAG_Query)
                .then(res =>
                    (setTagsOption(res.data.map(item =>({id: counter++, name: item.tag})))))
                .catch(e => console.log(e))
        }

        const fetchGroups = async () => {
            let counter = 0
            await fetchData(GROUPS_URL)
                .then(res =>
                    (setgroupsOption(res.data.map(item =>({id: counter++, name: item.name})))))
                .catch(e => console.log(e))
        }

        fetchTags().catch(e => console.log(e))
        fetchGroups().catch(e => console.log(e))
    }, [])


    const handleSubmit = async () => {
        setIsPending(true)
        await postModel(INCIDENT_URL, stateTest)
            .then(() => setIsPending(false))
            .catch((err) => console.log(err))
    }

    console.log(groupsOption)

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
        <div className={"incidentCreate"}>
            <form onSubmit={handleSubmit} onKeyDown={(e) => { e.key === 'Enter' && e.preventDefault(); }}>
            <h2>New Incident</h2>

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
                {tagsOption.length > 0  ?  <TagsInput
                    setTagsFunc={setTags}
                    data={tagsOption}
                    /> : null }
            </label>

                {groupsOption.length > 0 ?
                    <GroupSelectComponent
                        data={groupsOption}
                        setSelectedFunc={setSelectedGroups}
                    />: null
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
            {!isPending && <button className={"bnt"}>SEND INCIDENT</button>}
            {isPending && <button className={"bnt"} disabled>Sending mail...</button>}
            </form>
        </div>
    )
}

export default Incident;
