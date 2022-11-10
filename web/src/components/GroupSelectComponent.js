import "./GroupSelectComponent.css";
import {useEffect, useState} from "react";
import "../constants/DummyGroups.js"
import Select from "react-select"
import {grey} from "@mui/material/colors";
import customStyles from "./SelectStyle";


//TODO if changed make select to isMulti
function GroupSelectComponent({ data, setSelectedFunc }) {
    const [warning, setWarning] = useState("")

    const [options, setOptions] = useState([])

    useEffect(() => {
        setOptions(data.map(item => ({value: item.id, label: item.name})))
    }, [])

    function handleChangeOptions(e) {
        /*const group = (e.map(item => item.value).join(","))
        setWarning({
            ...warning, "receiverGroup": group
        })*/
        const group = e.label
        setSelectedFunc(group)
    }



    return (
        <div>
            <label>Receiver Group:
                <Select styles={customStyles} options={options} onChange={handleChangeOptions} className={"input-group select-group"}/>
            </label>

        </div>
    )
}

export default GroupSelectComponent
