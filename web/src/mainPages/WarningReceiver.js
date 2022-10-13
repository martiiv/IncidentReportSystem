import Table from "../components/Table";
import Persons from "../constants/DummyPersons";
import RowWR from "../constants/RowWR";
import {Component} from "react";


class WarningReceiver extends Component {
    state = {
        value: []
    };

    /**
     * Function that will get store the id of the checked persons in an array.
     * If unchecked the id will be removed from the array.
     * @param val person to be checked or unchecked.
     */
    onChangeValueHandler = val => {
        if (val.target.checked){
            this.setState(prevState =>( {
                    value: [...prevState.value, val.target.value]
                })
            )
        }else{
            let filteredArray = this.state.value.filter(item => item !== val.target.value)
            this.setState({value: filteredArray});
        }
    };


    render() {
        const {value} = this.state;
        return (
                <div>
                    <h1>Warning Receiver</h1>
                    <Table
                        type={"WR"}
                        data={Persons}
                        row={RowWR}
                        value={value}
                        onChangeValue={this.onChangeValueHandler}
                    />
                    <div className={"button-group"} style={{display: "flex", justifyContent: "center"}}>
                        <button className={"btn"}>Delete</button>
                        <button className={"btn"}>Create</button>
                        <button className={"btn"}>Group</button>
                    </div>
                </div>
        );
    }
}

export default WarningReceiver;



