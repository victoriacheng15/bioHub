export const NAME = "Victoria Cheng";
export const HEADLINE =
	"Software Developer | React | TypeScript | Python | Docker | Node.js | MongoDB";


const LIST = [
    {
      title: "GitHub",
      href: "https://github.com/victoriacheng15",
      path: "./icons/github.svg",
      category: "social",
    },
    {
      title: "LinkedIn",
      href: "https://www.linkedin.com/in/victoriacheng15/",
      path: "./icons/linkedin.svg",
      category: "social",
    },
    {
      title: "TikTok",
      href: "https://www.tiktok.com/@victoriacheng15",
      path: "./icons/tiktok.svg",
      category: "social",
    },
    {
      title: "youtube",
      href: "https://www.youtube.com/@victoriacheng15",
      path: "./icons/youtube.svg",
      category: "social",
    },
    {
      title: "Twitter",
      href: "https://twitter.com/viktoriacheng15",
      path: "./icons/twitter.svg",
      category: "social",
    },
    {
      title: "Buy me a coffee",
      href: "https://www.buymeacoffee.com/victoriacheng15",
      path: "./icons/buymeacoffee.svg",
      category: "social",
    },
    {
      title: "🐧 Portfolio site 🐧",
      href: "https://victoriacheng15.vercel.app",
    },
    {
      title: "React Hooks Series",
      href: "https://victoriacheng15.vercel.app/tags/react",
    },
  ];
  
  export const SCOIALS = LIST.filter(({ category }) => category);
  export const LINKS = LIST.filter(({ category }) => !category);