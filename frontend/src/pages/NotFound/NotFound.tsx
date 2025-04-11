import styles from "./NotFound.module.css";

export default function NotFound() {
  return (
    <div className={styles.notFound}>
      <h1>404 - Page Not Found</h1>
      <p>The page you are looking for does not exist.</p>
    </div>
  );
}
