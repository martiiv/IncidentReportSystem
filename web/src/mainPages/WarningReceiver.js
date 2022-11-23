import Table from "../components/Table";
import RowWR from "../constants/RowWR";
import {Component} from "react";
import "./WarningReceiver.css"
import fetchData from "../middleware/FetchData";
import {RECEIVER_URL} from "../constants/WebURL";
import deleteData from "../middleware/deleteData";
import {Link} from "react-router-dom";
import {ReactNotifications, Store} from "react-notifications-component";


/**
 * Class warningReceiver, will display a table with warning receivers.
 */
class WarningReceiver extends Component {
    state = {
        value: [],
        data: []
    };

    /**
     * Function that will get store the id of the checked persons in an array.
     * If unchecked the id will be removed from the array.
     * @param val person to be checked or unchecked.
     */
    onChangeValueHandler = val => {
        const { checked } = val.target;
        console.log(val.target.value)
        if (checked) {
            this.setState(prevState => ({
                    value: [...prevState.value, {id: (val.target.value)}]
                })
            )
        }else{
            let filteredArray = this.state.value.filter(item =>  item.id !== val.target.value)
            this.setState({value: filteredArray});
            console.log(this.state.value)
        }
    };




    /**
     * Function that will run when the class first is initialized.
     * Function will fetch data from the API, and store it in the data array.
     * This is an async function, so this will be run as a new thread.
     * @returns {Promise<void>} the response of the promise.
     */
    async componentDidMount() {
        await fetchData(RECEIVER_URL)
            .then((warningReceiver) => {
                this.setState({data: warningReceiver.data})
            })
            .catch(() => alert("No connection, try again later"))
    }


    /**
     * Function that will delete warning receivers.
     * The function will call on the API, and changes will happen in the database.
     * When the selected receivers is deleted, the user will get a notification on the screen.
     * If the API call does fail, the user will get an alert.
     *
     * @returns {Promise<void>} the status of the promise. Either rejection or resolve.
     */
    deleteWarningReceivers = async () => {
        const {value} = this.state;
        // eslint-disable-next-line no-restricted-globals
        if (confirm('Are you sure you want to save this thing into the database?')) {
            await deleteData(RECEIVER_URL, value)
                .then(() => {
                    Store.addNotification({
                        title: "Successfully Deleted!",
                        message: "Warning Receiver is successfully deleted",
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
                    window.location.reload(true)
                }, 2000))
                .catch(() => {
                    alert("something went wrong. please try again")
                })

        }
    }



    render() {
        const {value, data} = this.state;
        return (
            <div>
                <h1>Warning Receiver</h1>
                <ReactNotifications />

                <Table
                    type={"WR"}
                    data={data}
                    row={RowWR}
                    value={value}
                    onChangeValue={this.onChangeValueHandler}
                />
                <div className={"button-group"} style={{display: "flex", justifyContent: "center"}}>
                    <button className={"btn btnGroup"} onClick={this.deleteWarningReceivers}>Delete</button>
                    <Link to={RECEIVER_URL + "/new"}>
                        <button className={"btn btnGroup"}>Create</button>
                    </Link>

                    <Link to={RECEIVER_URL + "/group"}>
                        <button className={"btn btnGroup"}>Group</button>
                    </Link>
                </div>
            </div>
        );
    }
}

export default WarningReceiver;



