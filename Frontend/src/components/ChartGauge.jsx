export default function SemiCircularGauge({ value, min = 0, max = 100, centerText }) {
  const radius = 117;
  const strokeWidth = 26;
  const circumference = Math.PI * radius;
  const percent = Math.min(Math.max((value - min) / (max - min), 0), 1);
  const offset = circumference * (1 - percent);

  return (
    <svg width={286} height={169} viewBox="0 0 286 169">
      <defs>
        <linearGradient id="gaugeGradient" x1="0%" y1="0%" x2="100%" y2="0%">
          <stop offset="0%" stopColor="#4caf50" /> {/* verde moderato */}
          <stop offset="50%" stopColor="#ffeb3b" /> {/* giallo tenue */}
          <stop offset="100%" stopColor="#f44336" /> {/* rosso professionale */}
        </linearGradient>
      </defs>

      <path
        d={`
          M 26 143
          A ${radius} ${radius} 0 0 1 260 143
        `}
        fill="none"
        stroke="#3d3d3d94"
        strokeWidth={strokeWidth}
        strokeLinecap="round"
      />

      <path
        d={`
          M 26 143
          A ${radius} ${radius} 0 0 1 260 143
        `}
        fill="none"
        stroke="url(#gaugeGradient)"
        strokeWidth={strokeWidth}
        strokeLinecap="round"
        strokeDasharray={circumference}
        strokeDashoffset={offset}
      />

      <text
        x="143"
        y="130"
        textAnchor="middle"
        fontSize="33"
        fill="white"
        fontWeight="bold"
        fontFamily="system-ui, Avenir, Helvetica, Arial, sans-serif"
      >
        {centerText ?? `${Math.round(value)}`}
      </text>

      <text
        x="143"
        y="156"
        textAnchor="middle"
        fontSize="20"
        fill="#3d3d3d94"
        fontWeight="bold"
        fontFamily="system-ui, Avenir, Helvetica, Arial, sans-serif"
      >
        value: {value ?? undefined}
      </text>
    </svg>
  );
}
