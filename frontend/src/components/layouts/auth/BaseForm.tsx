import style from "./BaseForm.module.css";

interface InputField extends React.InputHTMLAttributes<HTMLInputElement> {}

export function InputBox(input: InputField) {
  return (
    <div className={style.input_box}>
      <input {...input} />
    </div>
  );
}

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  text: string;
}

export interface BaseFormProps {
  title?: string;
  inputs: InputField[];
  button: ButtonProps;
}

export default function BaseFrom(props: BaseFormProps) {
  return (
    <div className={style.wrapper}>
      <form className={style.form}>
        <h1>{props.title}</h1>
        <div className={style.input_container}>
          {props.inputs.map((input, index) => (
            <InputBox key={index} {...input} />
          ))}
        </div>
        <button {...props.button} className={style.button}>
          {props.button.text}
        </button>
      </form>
    </div>
  );
}
