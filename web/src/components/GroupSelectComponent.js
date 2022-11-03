import "./GroupSelectComponent.css";
import {useEffect, useState} from "react";
import "../constants/DummyGroups.js"
import Select from "react-select"

function GroupSelectComponent({ data, setSelectedFunc }) {
    const [warning, setWarning] = useState([])

    const [options, setOptions] = useState([])

    useEffect(() => {
        setOptions(data.map(item => ({value: item.id, label: item.name})))
    }, [])

    function handleChangeOptions(e) {
        const group = (e.map(item => item.value).join(","))
        setWarning({
            ...warning, "receiverGroup": group
        })
        setSelectedFunc(group)
    }

    return (
        <div>
            <label>Receiver Group:
                <Select isMulti options={options} onChange={handleChangeOptions} className={"input-group"}/>
            </label>

        </div>
    )
}

export default GroupSelectComponent
