import mstyle from "./ContributionCard.module.css";

export interface ContributionCardProps {
  img: string;
  text: string;
  style?: React.CSSProperties;
}

export default function ContributionCard({
  img,
  text,
  style,
}: ContributionCardProps) {
  return (
    <div
      className={mstyle.container}
      style={{
        backgroundImage: `url(${img})`,
        backgroundRepeat: "no-repeat",
        backgroundPosition: "right",
        ...style,
      }}
    >
      <p>{text}</p>
    </div>
  );
}
