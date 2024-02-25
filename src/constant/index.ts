export const NAME = "Victoria Cheng";
export const HEADLINE =
	"🌊 Software Developer | React | TypeScript | Node.js | MongoDB | Docker | JavaScript";


const LIST = [
    {
      title: "GitHub",
      href: "https://github.com/victoriacheng15",
      category: "social",
    },
    {
      title: "LinkedIn",
      href: "https://www.linkedin.com/in/victoriacheng15/",
      category: "social",
    },
    {
      title: "Twitter",
      href: "https://twitter.com/viktoriacheng15",
      category: "social",
    },
    {
      title: "YouTube",
      href: "https://www.youtube.com/@victoriacheng15",
      category: "social",
    },
    {
      title: "Buy me a coffee",
      href: "https://www.buymeacoffee.com/victoriacheng15",
      category: "social",
    },
    {
      title: "👋 🐧 Portfolio site 🐧 👋",
      href: "https://victoriacheng15.vercel.app",
    },
    {
      title: "React Hooks Series",
      href: "https://victoriacheng15.vercel.app/tags/react-hooks",
    },
  ];
  
  export const SCOIALS = LIST.filter(({ category }) => category);
  export const LINKS = LIST.filter(({ category }) => !category);