package bd

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	IpsCollection        *mongo.Collection
	PersonasIpsColletion *mongo.Collection
	Ctx                  = context.TODO()
)

/*Setup opens a database connection to mongodb*/
func Setup() {
	host := "127.0.0.1"
	port := "27017"
	connectionURI := "mongodb://" + host + ":" + port + "/"
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("scrapeo")
	IpsCollection = db.Collection("ips")
	PersonasIpsColletion = db.Collection("personas_ips")
}

type Persona struct {
	PEIPS_NRO_DOCUMENTO   string `bson:"pers_codigo"`
	PEIPS_NOMBRES         string `bson:"pers_nombre"`
	PEIPS_APELLIDOS       string `bson:"pers_apellidos"`
	PEIPS_FEC_NACIMIENTO  string `bson:"pers_fec_nacimiento"`
	PEIPS_SEXO            string `bson:"pers_sexo"`
	PEIPS_TIPO_ASEG       string `bson:"pers_tipo_aseg"`
	PEIPS_BENEFICIARIOS   string `bson:"pers_beneficiarios"`
	PEIPS_ENROLADO        string `bson:"pers_enrolado"`
	PEIPS_VENC_FE_VIDA    string `bson:"pers_venc_fe_vida"`
	PEIPS_NRO_PATRONAL    string `bson:"pers_nro_patronal"`
	PEIPS_EMPLEADOR       string `bson:"pers_empleador"`
	PEIPS_ESTADO_IPS      string `bson:"pers_estado_ips"`
	PEIPS_MESES_APORTE    string `bson:"pers_meses_aporte"`
	PEIPS_VENCIMIENTO     string `bson:"pers_vencimiento"`
	PEIPS_ULT_PER_ABONADO string `bson:"pers_ult_per_abonado"`
}

func CreatePersonasIps(b Persona) (string, error) {
	result, err := PersonasIpsColletion.InsertOne(Ctx, b)
	if err != nil {
		fmt.Println("ocurrió un error: ", err)
		return "0", err
	}
	fmt.Println("Insertó correctamente")
	return fmt.Sprintf("%v", result.InsertedID), err
}
func GetCedulas() ([]Persona, error) {
	var cedula Persona
	var cedulas []Persona

	cursor, err := IpsCollection.Find(Ctx, bson.M{"PEIPS_VERIFICADO": "N"})
	if err != nil {
		defer cursor.Close(Ctx)
		return cedulas, err
	}

	for cursor.Next(Ctx) {
		err := cursor.Decode(&cedula)
		if err != nil {
			return cedulas, err
		}
		cedulas = append(cedulas, cedula)
	}

	return cedulas, nil
}
func GetCi(id string) (Persona, error) {
	var b Persona
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return b, err
	}

	err = IpsCollection.
		FindOne(Ctx, bson.D{{"_id", objectId}}).
		Decode(&b)
	if err != nil {
		return b, err
	}
	return b, nil
}
func UpdatePersonaIPS(ci Persona, p_verificado string) error {
	filter := bson.D{{"PEIPS_NRO_DOCUMENTO", ci.PEIPS_NRO_DOCUMENTO}}
	update := bson.D{{"$set", bson.D{{"PEIPS_VERIFICADO", p_verificado}}}}
	_, err := IpsCollection.UpdateOne(
		Ctx,
		filter,
		update,
	)
	fmt.Println("Se actualizó correctamente")
	return err
}
