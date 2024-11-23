import BaseFrom, { BaseFormProps } from "./BaseForm";

const loginData: BaseFormProps = {
  title: "Login",
  inputs: [
    { required: true, type: "email", placeholder: "Input email" },
    { required: true, type: "password", placeholder: "Input password" },
  ],
  button: {
    text: "Войти",
  },
};
export default function Login() {
  return <BaseFrom {...loginData} />;
}
