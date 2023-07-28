import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { registerRoute } from "../utils/APIRoutes";
import axios from "axios";
import { z } from "zod";

const useCustomRegister = () => {
  const schema = z
    .object({
      username: z
        .string()
        .min(4, { message: "The username must be 4 characters or more" })
        .max(10, { message: "The username must be 10 characters or less" })
        .regex(/^[a-zA-Z0-9_]+$/, {
          message:
            "The username must contain only letters, numbers and underscore (_)",
        }),
      email: z
        .string()
        .email({ message: "Invalid email" })
        .min(1, { message: "Email is required" }),
      password: z
        .string()
        .min(1, { message: "Password is required" })
        .min(8, { message: "Password must have more than 8 characters" }),
      confirmPassword: z
        .string()
        .min(1, { message: "Password confirmation is required" }),
    })
    .refine((data) => data.password === data.confirmPassword, {
      path: ["confirmPassword"],
      message: "Passwords do not match",
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
      const { username, email, password, confirmPassword } = values;
      const { data: response } = await axios.post(registerRoute, {
        username,
        email,
        password,
        confirmPassword,
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
export default useCustomRegister;
