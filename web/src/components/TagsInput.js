// src/components/TagsInput.js
import './TagsInput.css';
import {useEffect, useState} from 'react'
import CreatableSelect from 'react-select/creatable';

/**
 * TODO: Add API connection after Tags-table is created in the DB
 * TODO: Add write support to API when new tags are created
 * Code is mostly taken from https://dev.to/0shuvo0/lets-create-an-add-tags-input-with-react-js-d29 as this
 * example adds tags in a way we want to.
 * @returns {JSX.Element}
 * @constructor
 */
function TagsInput({setTagsFunc, data}){
    const [tags, setTags] = useState([])

    const [options, setOptions] = useState([])

    useEffect(() => {
        setOptions(data.map(item => ({value: item.id, label: item.name})))
    }, [tags])

    function handleChangeOptions(e) {
        const tag = (e.map(item => item.value).join(","))
        setTags({
            ...tags, "receiverGroup": tag
        })
        setTagsFunc(tag)
    }

    return (
        <div className={"tag-container"}>
            <CreatableSelect isMulti options={options} onChange={handleChangeOptions} className={"input-group"}/>
        </div>
    )
}

export default TagsInput
