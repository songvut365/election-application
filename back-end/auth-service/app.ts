import express, { Application, Request, Response, urlencoded } from "express";

const app: Application = express();
const PORT = 5001;

app.use(express.json());
app.use(urlencoded({ extended: true }));

app.get("/", async(req: Request, res: Response): Promise<Response> => {
  return res.status(200).send({
    message: "Auth Service"
  });
});

try {
  app.listen(PORT, ():void => {
    console.log(`Connected successfully on port http://localhost:${PORT}`);
  });
} catch (err: any) {
  console.log(`Error occurred ${err.message}`);
}