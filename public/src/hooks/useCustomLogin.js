import { useState, useEffect } from "react";
import { z } from "zod";
import axios from "axios";
import { loginRoute } from "../utils/APIRoutes";
import { useNavigate } from "react-router-dom";

const useCustomLogin = () => {
  const schema = z.object({
    username: z
      .string()
      .min(4, { message: "The username must be 4 characters or more" })
      .max(10, { message: "The username must be 10 characters or less" })
      .regex(/^[a-zA-Z0-9_]+$/, {
        message:
          "The username must contain only letters, numbers and underscore (_)",
      }),
    password: z
      .string()
      .min(1, { message: "Password is required" })
      .min(8, { message: "Password must have more than 8 characters" }),
  });

  const [errorMessage, setErrorMessage] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    if (localStorage.getItem(process.env.REACT_APP_LOCALHOST_KEY)) {
      navigate("/");
    }
  }, []);

  const onSubmit = async (data) => {
    try {
      const values = schema.parse(data);
      const { username, password } = values;
      const { data: response } = await axios.post(loginRoute, {
        username,
        password,
      });
      if (response.status === false) {
        setErrorMessage(response.msg);
      }
      if (response.status === true) {
        localStorage.setItem(
          process.env.REACT_APP_LOCALHOST_KEY,
          JSON.stringify(response.user)
        );
        navigate("/");
      }
    } catch (error) {
      console.error(error);
    }
  };

  return { schema, errorMessage, onSubmit };
};
export default useCustomLogin;
