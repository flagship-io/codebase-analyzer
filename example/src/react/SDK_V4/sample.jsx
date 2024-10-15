import React from "react";
import { useFsFlag } from "@flagship.io/react-sdk";

export const MyReactComponent = () => {
    ///Step 2:  Retrieves a flag named "backgroundColor"
    const flag = useFsFlag("backgroundColor")

    //Step 3: Returns the value of the flag or if the flag does not exist, it returns the default value "green" 
    const flagValue = flag.getValue("green")


    const flag_ = useFsFlag("btnSize").getValue(16)

    const flag1 = useFsFlag("showBtn")
    const flagValue_ = flag1.getValue(true)
    
    return (
        <button
            style={{
                height: "200px",
                width: "200px",
                backgroundColor: flagValue,
            }}
        >
            {"I'm a square with color=" + flagValue}
        </button>
    );
};