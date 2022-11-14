import React from "react";
import { Navigate } from "react-router-dom";

/**
 * Function that will check if the user is authenticated, before sending the user to a protected route
 *
 * @param children the element route.
 * @returns {JSX.Element|*}
 */
const ProtectedRoute = ({ children }) => {
    const credentials = (sessionStorage.getItem("credentials"))


    console.log("Check user in Private: ", credentials);


    if (!credentials) {
        return <Navigate to={"/"} />;
    }
    return children;
};

export default ProtectedRoute;
