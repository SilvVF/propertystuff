import { useState } from "react";
import { Greet, Atlas } from "../wailsjs/go/main/App";
import { Button } from "./components/ui/button";
import { Input } from "./components/ui/input";
import { SearchIcon } from "lucide-react";
import { Separator } from "./components/ui/separator";
import { RadioGroup, RadioGroupItem } from "./components/ui/radio-group";
import { Label } from "./components/ui/label";

function App() {
  const [resultText, setResultText] = useState("");

  function atlas(query: string) {
    Atlas(query).then();
  }

  const [selected, setSelected] = useState("p")

  return (
    <div className="flex flex-col h-screen w-screen items-center justify-center space-y-4">
      <RadioGroup value={selected} onValueChange={(v) => setSelected(v)}>
        <div className="flex items-center space-x-2">
          <RadioGroupItem value="p" id="p" />
          <Label htmlFor="option-one">Philadelphia</Label>
        </div>
        <div className="flex items-center space-x-2">
          <RadioGroupItem value="d" id="d" />
          <Label htmlFor="d">Delaware</Label>
        </div>
      </RadioGroup>
      <SearchBar onSubmit={(q) => atlas("2051 W OXFORD ST")}></SearchBar>
      <div>
        {resultText}
      </div>
    </div>
  );
}

function SearchBar({ onSubmit }: { onSubmit: (query: string) => void }) {
  const handleSubmit = (e: any) => {
    e.preventDefault();
    onSubmit(e.target.value);
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="flex flex-row rounded-full h-fit bg-primary/30 items-center space-x-1 w-2/5"
    >
      <Input
        className="font-semibold text-foreground bg-transparent border-none outline-none focus:outline-none rounded-full"
        placeholder="Search..."
      />
      <Separator orientation="vertical" className="h-6 bg-foreground" />
      <Button
        className="bg-transparent hover:bg-transparent hover:outline-none pe-4"
        size="icon"
        type="submit"
        variant="link"
      >
        <SearchIcon />
      </Button>
    </form>
  );
}

export default App;
