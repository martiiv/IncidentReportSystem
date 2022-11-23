import "./Profile.css"

/**
 * Display the profile page
 * @returns {JSX.Element}
 * @constructor
 */
function Profile() {

    /**
     * Fetch the credentials data.
     * @type {any}
     */
    const credentials = JSON.parse(sessionStorage.getItem("credentials"))

    return (
        <div className="Profile">
            <div className={"card"}>
                <p>Name: {credentials.userName}</p>
                <p>Company: {credentials.company}</p>
                <p>Email: {credentials.email}</p>
                <p>ID: {credentials.cid}</p>
            </div>
            <button className={"btn"} onClick={() => {
                sessionStorage.clear()
                window.location.reload(true);
            }}>Log out
            </button>
        </div>

    );


}

export default Profile;
