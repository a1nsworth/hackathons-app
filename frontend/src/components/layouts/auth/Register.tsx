import BaseFrom, { BaseFormProps } from "./BaseForm";

const loginData: BaseFormProps = {
  title: "Register",
  inputs: [
    { required: true, type: "email", placeholder: "Input email" },
    { required: true, type: "password", placeholder: "Input password" },
    { required: true, type: "password", placeholder: "Confirm password" },
  ],
  button: {
    text: "Зарестрироваться",
  },
};

export default function Register() {
  return <BaseFrom {...loginData} />;
}
