import "./GroupSelectComponent.css";
import {useEffect, useState} from "react";
import Select from "react-select"
import customStyles from "./SelectStyle";


/**
 * Function that will create a select component.
 * @param data the dataset that will be selectable by the user.
 * @param setSelectedFunc Function that will set the group selected by the user.
 * @returns {JSX.Element} The Selector with label.
 * @constructor
 */
function GroupSelectComponent({ data, setStateTestFunc }) {
    const [options, setOptions] = useState([])


    useEffect(() => {
        setOptions(data.map(item => ({value: item.id, label: item.name})))
    }, [])


    /**
     * Function that will be called when the value in Select is changed.
     *
     * @param e input from Select
     */
    function handleChangeOptions(e) {
        const group = e.label
        setStateTestFunc(prevState => ({
            ...prevState,
            "receivingGroup": group,
        }))
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
