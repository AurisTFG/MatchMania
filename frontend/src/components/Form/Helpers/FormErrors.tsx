import styles from './FormErrors.module.css';

export type FormErrorsProps = {
  messages: string[];
  marginTop?: string | number;
};

export default function FormErrors({ messages, marginTop }: FormErrorsProps) {
  if (messages.length === 0) {
    return null;
  }

  return (
    <div
      className={styles.errorContainer}
      style={marginTop ? { marginTop } : undefined}
    >
      {messages.map((message, index) => (
        <div
          key={index}
          className={styles.errorMessage}
        >
          {message}
        </div>
      ))}
    </div>
  );
}
