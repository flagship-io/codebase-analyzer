import React from "react";
import { useFsModifications } from "@flagship.io/react-sdk";

// fs:flag:btnColor
const btnColorFlag = 'btnColor';
const fsModificationsDynamic = useFsModifications([
    {
      key: btnColorFlag,
      defaultValue: "red",
      activate: false,
    },
  ]);

export const MyReactComponent = () => {
  const fsModifications = useFsModifications([
    {
      key: "backgroundColor",
      defaultValue: "green",
      activate: false,
    },
  ]);
  return (
    <div
      style={{
        height: "200px",
        width: "200px",
        backgroundColor: fsModifications.backgroundColor,
      }}
    >
      {"I'm a square with color=" + fsModifications.backgroundColor}
    </div>
  );
};
