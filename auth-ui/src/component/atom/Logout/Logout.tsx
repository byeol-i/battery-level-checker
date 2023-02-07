import React from "react";

const Logout = ({onLogOut}:any) => {
   
    return (
        <div>
            {onLogOut && <button onClick={onLogOut}>Logout</button>}
        </div>
    )
}

export default Logout