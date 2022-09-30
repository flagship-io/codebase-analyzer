import React from "react";
import {useFsFlag} from "@flagship.io/react-sdk";

export const MyReactComponent = () => {
    const backgroundColorFlag = useFsFlag("backgroundColor", "green")
    const btnColorFlag = useFsFlag("btnColor", "red")

    return (
        <div
            style={{
                height: "200px",
                width: "200px",
                backgroundColor: backgroundColorFlag.getValue(),
            }}
        >
            {"I'm a square with color=" + backgroundColorFlag.getValue()}
            <button style={{color: btnColorFlag.getValue()}}>Button</button>
        </div>
    );
};